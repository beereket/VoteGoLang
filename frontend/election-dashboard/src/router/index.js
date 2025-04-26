import Vue from 'vue';
import Router from 'vue-router';
// import { jwtDecode } from 'jwt-decode';
import OnlineDashboard from "@/views/OnlineDashboard.vue";
import LoginPage from "@/views/LoginPage.vue";


Vue.use(Router);

const router = new Router({
    mode: 'history',
    routes: [
        { path: '/login', component: LoginPage },
        {
            path: '/admin/dashboard',
            component: OnlineDashboard,
            meta: { requiresAdmin: true },
        },
    ]
});

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token');
    const role = localStorage.getItem('role');

    if (to.matched.some(record => record.meta.requiresAdmin)) {
        if (!token || role !== 'admin') {
            return next('/login');
        }
    }

    next();
});

export default router;
