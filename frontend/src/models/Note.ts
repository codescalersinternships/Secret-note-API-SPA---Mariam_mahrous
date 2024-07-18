export interface Note {
  id: number
  title: string
  content: string
  current_views: number
  max_views: number
}

export class NoteData {
  title: string
  content: string
  expiration_date?: string
  max_views?: string

  constructor(title: string, content: string, expiration_date?: string, max_views?: string) {
    this.title = title
    this.content = content
    this.expiration_date = expiration_date
    this.max_views = max_views
  }
}
