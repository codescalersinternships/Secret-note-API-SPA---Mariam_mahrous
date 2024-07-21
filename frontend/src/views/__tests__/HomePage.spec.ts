import { describe, vi, expect, test, type Mock } from 'vitest'
import { mount } from '@vue/test-utils'
import HomePage from '../HomePage.vue'

global.fetch = vi.fn(() =>
  Promise.resolve({
    ok: true,
    json: () => Promise.resolve({ token: 'test token' })
  })
) as Mock

describe('Homepage view', () => {
  const wrapper = mount(HomePage)

  test('renders the correct components', () => {
    expect(wrapper.findAll('button[type="submit"]').length).toBe(2)
    expect(wrapper.find('#view-Notes').text()).toBe('View Notes')
    expect(wrapper.find('#new-note').text()).toBe('Create a new note')
    expect(wrapper.find('h1').text()).toBe('Homepage')
  })
})
