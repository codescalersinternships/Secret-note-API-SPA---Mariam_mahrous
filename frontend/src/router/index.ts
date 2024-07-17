import { createRouter, createWebHistory } from 'vue-router'
import SignUp from '../views/SignUp.vue'
import CreateNote from '../views/CreateNote.vue'
import ViewNote from '../views/ViewNote.vue'
import ViewNotes from '../views/ViewNotes.vue'
import HomePage from '../views/HomePage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', component: SignUp },
    { path: '/note/create', component: CreateNote },
    { path: '/note/:uuid', component: ViewNote, props: (route) => ({ noteID: route.params.uuid }) },
    { path: '/note', component: ViewNotes },
    { path: '/home', component: HomePage }
  ]
})

export default router
