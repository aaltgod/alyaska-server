import Home from './components/Home.vue'
import Archive from './components/Archive.vue'
import Upload from './components/Upload.vue'


export const routes = [
    {
        path: '/',
        component: Home,
    },
    {
        path: "/archive",
        component: Archive,
    },
    {
        path: "/upload",
        component: Upload,
    },
]