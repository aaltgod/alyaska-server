import Home from './components/Home.vue'
import Archive from './components/Archive.vue'


export const routes = [
    {
        path: '/',
        component: Home
    },
    {
        path: "/archive",
        component: Archive,
    }   
]