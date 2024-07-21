import { describe, vi, expect, test, type Mock } from 'vitest'
import { mount } from '@vue/test-utils'
import SignUp from '../SignUp.vue'

global.fetch = vi.fn(() =>
  Promise.resolve({
    ok: true,
    json: () => Promise.resolve({ token: 'test token' })
  })
) as Mock

describe('SignUp/LogIn view', () => {
  const wrapper = mount(SignUp)
  const API_URL = import.meta.env.VITE_API_URL

  test('renders the correct components', () => {
    expect(wrapper.find('input[type="email"]').exists()).toBe(true)
    expect(wrapper.find('input[type="password"]').exists()).toBe(true)
    expect(wrapper.find('button[type="submit"]').exists()).toBe(true)
    expect(wrapper.find('span').exists()).toBe(true)
    expect(wrapper.find('h1').exists()).toBe(true)
    expect(wrapper.find('label[for="Email"]').exists()).toBe(true)
    expect(wrapper.find('label[for="Password"]').exists()).toBe(true)
  })

  test('Page toggles correctly', async () => {
    expect(wrapper.html()).toContain('Sign up')
    expect(wrapper.html()).toContain('Already have an account?')
    await wrapper.find('span').trigger('click')
    expect(wrapper.html()).toContain('Sign up')
    expect(wrapper.html()).toContain('Create an account')
    await wrapper.find('span').trigger('click')
  })

  test('should call the api with correct data and navigate', async () => {
    wrapper.get('input[type="email"]').setValue('test@example.com')
    wrapper.get('input[type="password"]').setValue('password123')
    await wrapper.find('form').trigger('submit.prevent')

    expect(global.fetch).toHaveBeenCalledWith(
      API_URL + '/signup',
      expect.objectContaining({
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email: 'test@example.com', password: 'password123' })
      })
    )
  })
})
