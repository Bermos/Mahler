<template>
  <div
    class="canvas-container"
    @mousedown="startPan"
    @mousemove="pan"
    @mouseup="stopPan"
    @wheel="handleZoom"
  >
    <div class="canvas" :style="transformStyle">
      <!-- Dot grid background -->
      <div class="dot-grid" />

      <!-- Service cards -->
      <div
        v-for="(service, index) in services"
        :key="index"
        class="service-card"
        :style="{ left: service.x + 'px', top: service.y + 'px' }"
        @mousedown.stop="startDrag($event, index)"
      >
        <div class="service-icon" :class="service.type">
          <component :is="service.icon" />
        </div>
        <div class="service-content">
          <h3>{{ service.name }}</h3>
          <p class="service-url">{{ service.url }}</p>
          <div class="status">
            <div class="status-icon" />
            {{ service.status }}
          </div>
          <div v-if="service.replicas" class="replicas">{{ service.replicas }} Replicas</div>
        </div>
      </div>

      <!-- Connection lines -->
      <svg class="connections">
        <path
          v-for="(connection, index) in connections"
          :key="index"
          :d="calculatePath(connection)"
          class="connection-line"
        />
      </svg>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Box, Database, Server, Globe } from 'lucide-vue-next'

const scale = ref(1)
const position = ref({ x: 0, y: 0 })
const isDragging = ref(false)
const draggedService = ref(null)
const isPanning = ref(false)
const startPos = ref({ x: 0, y: 0 })

const services = ref([
  {
    name: 'frontend',
    type: 'js',
    url: 'frontend-prod.up.railway.app',
    status: 'Deployed just now',
    icon: Globe,
    x: 400,
    y: 200,
  },
  {
    name: 'backend',
    type: 'python',
    url: 'backend-prod.up.railway.app',
    status: 'Deployed just now',
    replicas: '3 Replicas',
    icon: Server,
    x: 700,
    y: 200,
  },
  {
    name: 'ackee analytics',
    type: 'analytics',
    url: 'ackee-prod.up.railway.app',
    status: 'Deployed via Docker Image',
    icon: Box,
    x: 200,
    y: 400,
  },
  {
    name: 'api gateway',
    type: 'api',
    url: 'api-prod.up.railway.app',
    status: 'Deployed just now',
    icon: Globe,
    x: 400,
    y: 500,
  },
  {
    name: 'postgres',
    type: 'db',
    url: 'pg-data',
    status: 'Deployed via Docker Image',
    icon: Database,
    x: 700,
    y: 500,
  },
])

const connections = ref([
  { from: 0, to: 3 },
  { from: 3, to: 1 },
  { from: 1, to: 4 },
  { from: 2, to: 3 },
])

const transformStyle = computed(() => ({
  transform: `translate(${position.value.x}px, ${position.value.y}px) scale(${scale.value})`,
}))

function startDrag(event, index) {
  isDragging.value = true
  draggedService.value = index
  startPos.value = {
    x: event.clientX - services.value[index].x,
    y: event.clientY - services.value[index].y,
  }
}

function startPan(event) {
  if (!isDragging.value) {
    isPanning.value = true
    startPos.value = {
      x: event.clientX - position.value.x,
      y: event.clientY - position.value.y,
    }
  }
}

function pan(event) {
  if (isDragging.value && draggedService.value !== null) {
    // Snap to grid (40px grid size)
    const snapToGrid = (value) => Math.round(value / 40) * 40
    services.value[draggedService.value].x = snapToGrid(event.clientX - startPos.value.x)
    services.value[draggedService.value].y = snapToGrid(event.clientY - startPos.value.y)
  } else if (isPanning.value) {
    position.value = {
      x: event.clientX - startPos.value.x,
      y: event.clientY - startPos.value.y,
    }
  }
}

function stopPan() {
  isDragging.value = false
  draggedService.value = null
  isPanning.value = false
}

function handleZoom(event) {
  event.preventDefault()
  const delta = event.deltaY > 0 ? 0.9 : 1.1
  scale.value = Math.min(Math.max(0.5, scale.value * delta), 2)
}

function calculatePath(connection) {
  const from = services.value[connection.from]
  const to = services.value[connection.to]

  // Card dimensions
  const cardWidth = 300
  const cardHeight = 160

  // Calculate center points
  const fromCenter = {
    x: from.x + cardWidth / 2,
    y: from.y + cardHeight / 2,
  }
  const toCenter = {
    x: to.x + cardWidth / 2,
    y: to.y + cardHeight / 2,
  }

  // Determine which sides to connect based on relative positions
  let startPoint, endPoint

  // Horizontal difference is greater than vertical
  if (Math.abs(toCenter.x - fromCenter.x) > Math.abs(toCenter.y - fromCenter.y)) {
    // From card is on the left
    if (fromCenter.x < toCenter.x) {
      startPoint = { x: from.x + cardWidth, y: fromCenter.y }
      endPoint = { x: to.x, y: toCenter.y }
    } else {
      startPoint = { x: from.x, y: fromCenter.y }
      endPoint = { x: to.x + cardWidth, y: toCenter.y }
    }
  } else {
    // From card is above
    if (fromCenter.y < toCenter.y) {
      startPoint = { x: fromCenter.x, y: from.y + cardHeight }
      endPoint = { x: toCenter.x, y: to.y }
    } else {
      startPoint = { x: fromCenter.x, y: from.y }
      endPoint = { x: toCenter.x, y: to.y + cardHeight }
    }
  }

  // Calculate middle points for the orthogonal path
  const midPoint = {
    x: startPoint.x + (endPoint.x - startPoint.x) / 2,
    y: startPoint.y + (endPoint.y - startPoint.y) / 2,
  }

  // Create path with right angles
  return `M ${startPoint.x} ${startPoint.y}
          L ${
            Math.abs(toCenter.x - fromCenter.x) > Math.abs(toCenter.y - fromCenter.y)
              ? midPoint.x
              : startPoint.x
          }
            ${
              Math.abs(toCenter.x - fromCenter.x) > Math.abs(toCenter.y - fromCenter.y)
                ? startPoint.y
                : midPoint.y
            }
          L ${
            Math.abs(toCenter.x - fromCenter.x) > Math.abs(toCenter.y - fromCenter.y)
              ? midPoint.x
              : endPoint.x
          }
            ${
              Math.abs(toCenter.x - fromCenter.x) > Math.abs(toCenter.y - fromCenter.y)
                ? endPoint.y
                : midPoint.y
            }
          L ${endPoint.x} ${endPoint.y}`
}
</script>

<style scoped>
.canvas-container {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  background-color: #0a0b14;
  position: relative;
}

.canvas {
  position: absolute;
  width: 100%;
  height: 100%;
  transform-origin: 0 0;
}

.dot-grid {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: radial-gradient(circle, rgba(255, 255, 255, 0.1) 1px, transparent 1px);
  background-size: 40px 40px;
}

.service-card {
  position: absolute;
  width: 300px;
  background: rgba(30, 32, 47, 0.9);
  border-radius: 8px;
  padding: 16px;
  cursor: move;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: white;
  transition: transform 0.2s;
}

.service-card:hover {
  transform: translateY(-2px);
}

.service-icon {
  width: 32px;
  height: 32px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 12px;
}

.service-icon.js {
  background: #f7df1e;
}
.service-icon.python {
  background: #3776ab;
}
.service-icon.analytics {
  background: #00bcd4;
}
.service-icon.api {
  background: #4caf50;
}
.service-icon.db {
  background: #2196f3;
}

.service-content h3 {
  margin: 0;
  font-size: 1.1rem;
  font-weight: 500;
}

.service-url {
  color: rgba(255, 255, 255, 0.6);
  font-size: 0.9rem;
  margin: 4px 0;
}

.status {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.9rem;
  color: rgba(255, 255, 255, 0.8);
  margin-top: 8px;
}

.status-icon {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #4caf50;
}

.replicas {
  margin-top: 8px;
  font-size: 0.9rem;
  color: rgba(255, 255, 255, 0.6);
}

.connections {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
}

.connection-line {
  stroke: rgba(255, 255, 255, 0.1);
  stroke-width: 2;
  stroke-dasharray: 4;
  fill: none;
  stroke-linecap: round;
  stroke-linejoin: round;
}
</style>
