import Home from './components/Home.vue'
import Archive from './components/Archive.vue'
import Upload from './components/Upload.vue'
import Folder from './components/Folder.vue'


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
    {
        path: "/folder/:folderName",
        component: Folder,
    }
]