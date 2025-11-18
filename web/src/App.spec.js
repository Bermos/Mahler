import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import App from './App.vue'
import ArchitectureCanvas from './components/ArchitectureCanvas.vue'

describe('App', () => {
  it('renders correctly', () => {
    const wrapper = mount(App)
    expect(wrapper.exists()).toBe(true)
  })

  it('renders ArchitectureCanvas component', () => {
    const wrapper = mount(App)
    expect(wrapper.findComponent(ArchitectureCanvas).exists()).toBe(true)
  })

  it('has no additional content besides ArchitectureCanvas', () => {
    const wrapper = mount(App)
    const canvas = wrapper.findComponent(ArchitectureCanvas)
    expect(canvas.exists()).toBe(true)
  })

  it('matches snapshot', () => {
    const wrapper = mount(App)
    expect(wrapper.html()).toMatchSnapshot()
  })
})
