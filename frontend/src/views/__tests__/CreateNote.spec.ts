import { describe, vi, expect, test, fail, type Mock } from 'vitest'
import { mount } from '@vue/test-utils'
import CreateNote from '../CreateNote.vue'
import type { title } from 'process'

global.fetch = vi.fn(() =>
  Promise.resolve({
    ok: true,
    json: () => Promise.resolve({ token: 'test token' })
  })
) as Mock

describe('CreateNote view', () => {
  const wrapper = mount(CreateNote)
  const API_URL = import.meta.env.VITE_API_URL

  test('renders the correct components', () => {
    expect(wrapper.find('label[for="title"]').exists()).toBe(true)
    expect(wrapper.find('input[type="text"]').exists()).toBe(true)
    expect(wrapper.find('label[for="content"]').exists()).toBe(true)
    expect(wrapper.find('textarea').exists()).toBe(true)
    expect(wrapper.find('label[for="content"]').exists()).toBe(true)
    expect(wrapper.find('input[type="date"]').exists()).toBe(true)
    expect(wrapper.find('label[for="max_views"]').exists()).toBe(true)
    expect(wrapper.find('input[type="number"]').exists()).toBe(true)
    expect(wrapper.find('button[type="submit"]').exists()).toBe(true)
  })

  test('Create a note test', async () => {
    wrapper.get('input[type="text"]').setValue('test')
    wrapper.get('textarea').setValue('test content')

    await wrapper.find('form').trigger('submit.prevent')

    expect(global.fetch).toHaveBeenCalledWith(
      API_URL + '/note/create',
      expect.objectContaining({
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer null'
        },
        body: JSON.stringify({ title: 'test', content: 'test content' })
      })
    )
  })
})
