import Vue from 'vue'
import VueRouter from 'vue-router'

import About from '@/pages/About.vue'
import Login from "@/pages/Login.vue"
import SmartFilter from "@/pages/SmartFilter.vue"
import SupplierPage from "@/pages/SupplierPage.vue"

Vue.use(VueRouter)

const routes = [{
        path: '/',
        name: 'login',
        component: Login
    },
    {
        path: '/about',
        name: 'About',
        component: About
    },
    {
        path: "/smart-filter",
        name: "SmartFilter",
        component: SmartFilter,
    },
    { path: '/supplier/:id', name: "SupplierPage", component: SupplierPage }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router