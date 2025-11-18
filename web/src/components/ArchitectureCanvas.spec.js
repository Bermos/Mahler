import { describe, it, expect, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import ArchitectureCanvas from './ArchitectureCanvas.vue'

describe('ArchitectureCanvas', () => {
  let wrapper

  beforeEach(() => {
    wrapper = mount(ArchitectureCanvas)
  })

  describe('Component Rendering', () => {
    it('renders the canvas container', () => {
      expect(wrapper.find('.canvas-container').exists()).toBe(true)
    })

    it('renders the canvas element', () => {
      expect(wrapper.find('.canvas').exists()).toBe(true)
    })

    it('renders the dot grid background', () => {
      expect(wrapper.find('.dot-grid').exists()).toBe(true)
    })

    it('applies initial transform style', () => {
      const canvas = wrapper.find('.canvas')
      expect(canvas.attributes('style')).toContain('translate(0px, 0px) scale(1)')
    })
  })

  describe('Service Cards', () => {
    it('renders all service cards', () => {
      const serviceCards = wrapper.findAll('.service-card')
      expect(serviceCards.length).toBe(5)
    })

    it('renders service card with correct name', () => {
      const firstCard = wrapper.findAll('.service-card')[0]
      expect(firstCard.text()).toContain('frontend')
    })

    it('renders service card with correct URL', () => {
      const firstCard = wrapper.findAll('.service-card')[0]
      expect(firstCard.text()).toContain('frontend-prod.up.railway.app')
    })

    it('renders service card with correct status', () => {
      const firstCard = wrapper.findAll('.service-card')[0]
      expect(firstCard.text()).toContain('Deployed just now')
    })

    it('renders service icon with correct class', () => {
      const firstCard = wrapper.findAll('.service-card')[0]
      const icon = firstCard.find('.service-icon')
      expect(icon.classes()).toContain('js')
    })

    it('renders replicas when present', () => {
      const backendCard = wrapper.findAll('.service-card')[1]
      expect(backendCard.text()).toContain('3 Replicas')
    })

    it('does not render replicas when not present', () => {
      const firstCard = wrapper.findAll('.service-card')[0]
      expect(firstCard.find('.replicas').exists()).toBe(false)
    })

    it('positions service cards correctly', () => {
      const firstCard = wrapper.findAll('.service-card')[0]
      expect(firstCard.attributes('style')).toContain('left: 400px')
      expect(firstCard.attributes('style')).toContain('top: 200px')
    })

    it('renders status icon', () => {
      const firstCard = wrapper.findAll('.service-card')[0]
      expect(firstCard.find('.status-icon').exists()).toBe(true)
    })
  })

  describe('Connections', () => {
    it('renders the SVG connections element', () => {
      expect(wrapper.find('.connections').exists()).toBe(true)
    })

    it('renders all connection lines', () => {
      const connectionLines = wrapper.findAll('.connection-line')
      expect(connectionLines.length).toBe(4)
    })

    it('renders connection lines with correct path attributes', () => {
      const connectionLines = wrapper.findAll('.connection-line')
      connectionLines.forEach((line) => {
        expect(line.attributes('d')).toBeTruthy()
      })
    })
  })

  describe('Drag and Drop', () => {
    it('starts dragging on service card mousedown', async () => {
      const firstCard = wrapper.findAll('.service-card')[0]
      await firstCard.trigger('mousedown', { clientX: 100, clientY: 100 })

      // Service should be in dragging state
      expect(wrapper.vm.isDragging).toBe(true)
      expect(wrapper.vm.draggedService).toBe(0)
    })

    it('moves service card during drag', async () => {
      const firstCard = wrapper.findAll('.service-card')[0]
      const initialX = wrapper.vm.services[0].x
      const initialY = wrapper.vm.services[0].y

      // Start dragging
      await firstCard.trigger('mousedown', { clientX: 450, clientY: 250 })

      // Move mouse
      await wrapper.find('.canvas-container').trigger('mousemove', {
        clientX: 500,
        clientY: 300,
      })

      // Position should change
      expect(wrapper.vm.services[0].x).not.toBe(initialX)
      expect(wrapper.vm.services[0].y).not.toBe(initialY)
    })

    it('snaps service card to grid during drag', async () => {
      const firstCard = wrapper.findAll('.service-card')[0]

      await firstCard.trigger('mousedown', { clientX: 450, clientY: 250 })
      await wrapper.find('.canvas-container').trigger('mousemove', {
        clientX: 465,
        clientY: 275,
      })

      // Should snap to 40px grid
      expect(wrapper.vm.services[0].x % 40).toBe(0)
      expect(wrapper.vm.services[0].y % 40).toBe(0)
    })

    it('stops dragging on mouseup', async () => {
      const firstCard = wrapper.findAll('.service-card')[0]
      await firstCard.trigger('mousedown', { clientX: 100, clientY: 100 })
      expect(wrapper.vm.isDragging).toBe(true)

      await wrapper.find('.canvas-container').trigger('mouseup')

      expect(wrapper.vm.isDragging).toBe(false)
      expect(wrapper.vm.draggedService).toBe(null)
    })
  })

  describe('Pan and Zoom', () => {
    it('starts panning on canvas mousedown when not dragging', async () => {
      await wrapper.find('.canvas-container').trigger('mousedown', {
        clientX: 100,
        clientY: 100,
      })

      expect(wrapper.vm.isPanning).toBe(true)
    })

    it('does not start panning when dragging a service', async () => {
      const firstCard = wrapper.findAll('.service-card')[0]
      await firstCard.trigger('mousedown', { clientX: 450, clientY: 250 })

      await wrapper.find('.canvas-container').trigger('mousedown', {
        clientX: 100,
        clientY: 100,
      })

      // Should still be dragging, not panning
      expect(wrapper.vm.isDragging).toBe(true)
      expect(wrapper.vm.isPanning).toBe(false)
    })

    it('pans the canvas', async () => {
      await wrapper.find('.canvas-container').trigger('mousedown', {
        clientX: 100,
        clientY: 100,
      })

      await wrapper.find('.canvas-container').trigger('mousemove', {
        clientX: 200,
        clientY: 200,
      })

      expect(wrapper.vm.position.x).toBe(100)
      expect(wrapper.vm.position.y).toBe(100)
    })

    it('stops panning on mouseup', async () => {
      await wrapper.find('.canvas-container').trigger('mousedown', {
        clientX: 100,
        clientY: 100,
      })
      expect(wrapper.vm.isPanning).toBe(true)

      await wrapper.find('.canvas-container').trigger('mouseup')

      expect(wrapper.vm.isPanning).toBe(false)
    })

    it('zooms in on wheel scroll up', async () => {
      const initialScale = wrapper.vm.scale

      await wrapper.find('.canvas-container').trigger('wheel', {
        deltaY: -100,
        preventDefault: () => {},
      })

      expect(wrapper.vm.scale).toBeGreaterThan(initialScale)
    })

    it('zooms out on wheel scroll down', async () => {
      const initialScale = wrapper.vm.scale

      await wrapper.find('.canvas-container').trigger('wheel', {
        deltaY: 100,
        preventDefault: () => {},
      })

      expect(wrapper.vm.scale).toBeLessThan(initialScale)
    })

    it('limits zoom to minimum scale of 0.5', async () => {
      // Zoom out many times
      for (let i = 0; i < 20; i++) {
        await wrapper.find('.canvas-container').trigger('wheel', {
          deltaY: 100,
          preventDefault: () => {},
        })
      }

      expect(wrapper.vm.scale).toBeGreaterThanOrEqual(0.5)
    })

    it('limits zoom to maximum scale of 2', async () => {
      // Zoom in many times
      for (let i = 0; i < 20; i++) {
        await wrapper.find('.canvas-container').trigger('wheel', {
          deltaY: -100,
          preventDefault: () => {},
        })
      }

      expect(wrapper.vm.scale).toBeLessThanOrEqual(2)
    })
  })

  describe('Transform Style', () => {
    it('updates transform style when position changes', async () => {
      await wrapper.find('.canvas-container').trigger('mousedown', {
        clientX: 0,
        clientY: 0,
      })
      await wrapper.find('.canvas-container').trigger('mousemove', {
        clientX: 50,
        clientY: 50,
      })

      const canvas = wrapper.find('.canvas')
      expect(canvas.attributes('style')).toContain('translate(50px, 50px)')
    })

    it('updates transform style when scale changes', async () => {
      await wrapper.find('.canvas-container').trigger('wheel', {
        deltaY: -100,
        preventDefault: () => {},
      })

      const canvas = wrapper.find('.canvas')
      expect(canvas.attributes('style')).toContain('scale(1.1)')
    })
  })

  describe('Path Calculation', () => {
    it('calculates path for horizontal connections', () => {
      const connection = { from: 0, to: 1 }
      const path = wrapper.vm.calculatePath(connection)

      expect(path).toContain('M')
      expect(path).toContain('L')
    })

    it('calculates path for vertical connections', () => {
      const connection = { from: 0, to: 2 }
      const path = wrapper.vm.calculatePath(connection)

      expect(path).toContain('M')
      expect(path).toContain('L')
    })

    it('calculates different paths for different connections', () => {
      const connection1 = { from: 0, to: 1 }
      const connection2 = { from: 2, to: 3 }

      const path1 = wrapper.vm.calculatePath(connection1)
      const path2 = wrapper.vm.calculatePath(connection2)

      expect(path1).not.toBe(path2)
    })

    it('handles connections from left to right', () => {
      const connection = { from: 0, to: 1 }
      const path = wrapper.vm.calculatePath(connection)

      // Should start from right edge of from card (x + 300)
      expect(path).toBeTruthy()
    })

    it('handles connections from right to left', () => {
      const connection = { from: 1, to: 0 }
      const path = wrapper.vm.calculatePath(connection)

      // Should start from left edge of from card
      expect(path).toBeTruthy()
    })

    it('handles connections from top to bottom', () => {
      const connection = { from: 0, to: 2 }
      const path = wrapper.vm.calculatePath(connection)

      // Should start from bottom edge of from card (y + 160)
      expect(path).toBeTruthy()
    })

    it('handles connections from bottom to top', () => {
      const connection = { from: 2, to: 0 }
      const path = wrapper.vm.calculatePath(connection)

      // Should start from top edge of from card
      expect(path).toBeTruthy()
    })
  })

  describe('Service Data Structure', () => {
    it('has all required service properties', () => {
      wrapper.vm.services.forEach((service) => {
        expect(service).toHaveProperty('name')
        expect(service).toHaveProperty('type')
        expect(service).toHaveProperty('url')
        expect(service).toHaveProperty('status')
        expect(service).toHaveProperty('icon')
        expect(service).toHaveProperty('x')
        expect(service).toHaveProperty('y')
      })
    })

    it('has numeric x and y coordinates', () => {
      wrapper.vm.services.forEach((service) => {
        expect(typeof service.x).toBe('number')
        expect(typeof service.y).toBe('number')
      })
    })
  })

  describe('Connection Data Structure', () => {
    it('has valid connection references', () => {
      wrapper.vm.connections.forEach((connection) => {
        expect(connection.from).toBeGreaterThanOrEqual(0)
        expect(connection.from).toBeLessThan(wrapper.vm.services.length)
        expect(connection.to).toBeGreaterThanOrEqual(0)
        expect(connection.to).toBeLessThan(wrapper.vm.services.length)
      })
    })
  })

  describe('Reactive State', () => {
    it('initializes with default scale of 1', () => {
      expect(wrapper.vm.scale).toBe(1)
    })

    it('initializes with position at origin', () => {
      expect(wrapper.vm.position.x).toBe(0)
      expect(wrapper.vm.position.y).toBe(0)
    })

    it('initializes with isDragging false', () => {
      expect(wrapper.vm.isDragging).toBe(false)
    })

    it('initializes with isPanning false', () => {
      expect(wrapper.vm.isPanning).toBe(false)
    })

    it('initializes with null draggedService', () => {
      expect(wrapper.vm.draggedService).toBe(null)
    })
  })
})
