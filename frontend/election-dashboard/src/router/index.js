import Vue from 'vue';
import Router from 'vue-router';
import LoginPage from "@/views/auth/LoginPage.vue";
import DashboardAnalytics from "@/views/admin/dashboards/DashboardAnalytics.vue";
import DashboardPage from "@/views/admin/dashboards/Dashboard.vue";
import NewsManagement from "@/views/admin/news/NewsManagement.vue";
import UserManagement from "@/views/admin/users/UserManagement.vue";
import ActivityLogs from "@/views/admin/ActivityLogs.vue";
import ElectionManagement from "@/views/admin/elections/ElectionManagement.vue";
import LiveResults from "@/views/public/LiveResults.vue";
import ElectionList from "@/views/public/elections/ElectionList.vue";
import ElectionPage from "@/views/public/elections/ElectionPage.vue";
import PetitionManagement from "@/views/admin/petitions/PetitionManagement.vue";
import PetitionList from "@/views/public/petitions/PetitionList.vue";
import PetitionDetail from "@/views/public/petitions/PetitionDetail.vue";
import RegisterPage from "@/views/auth/RegisterPage.vue";
import CandidateManagement from "@/views/admin/elections/CandidateManagement.vue";


Vue.use(Router);

const router = new Router({
    mode: 'history',
    routes: [

        { path: '/login', component: LoginPage, meta: { guestOnly: true } },
        { path: '/register', component: RegisterPage, meta: { guestOnly: true } },

        { path: '/', component: LiveResults },
        { path: '/elections', component: ElectionList },
        { path: '/elections/:id', component: ElectionPage, meta: { userOnly: true } },
        { path: '/petitions', component: PetitionList },
        { path: '/petitions/:id', component: PetitionDetail, meta: { userOnly: true } },

        {
            path: '/admin/dashboard',
            component: DashboardPage,
            meta: { requiresAdmin: true },
        },
        {
            path: '/admin/users',
            component: UserManagement,
            meta: { requiresAdmin: true },
        },
        {
            path: '/admin/analytics',
            component: DashboardAnalytics,
            meta: { requiresAdmin: true },
        },
        {
            path: '/admin/elections',
            component: ElectionManagement,
            meta: { requiresAdmin: true }
        },
        {
            path: '/admin/elections/:id/candidates',
            component: CandidateManagement,
            meta: { requiresAdmin: true },
        },
        {
            path: '/admin/petitions',
            component: PetitionManagement,
            meta: { requiresAdmin: true }
        },
        {
            path: '/admin/news',
            component: NewsManagement,
            meta: { requiresAdmin: true },
        },
        {
            path: '/admin/logs',
            component: ActivityLogs,
            meta: { requiresAdmin: true },
        },
    ]
});

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token');
    const role = localStorage.getItem('role');

    if (to.matched.some(record => record.meta.requiresAdmin)) {
        // Admin-only pages
        if (!token || role !== 'admin') {
            return next('/login');
        }
    }

    if (to.matched.some(record => record.meta.userOnly)) {
        // User-only pages
        if (!token) {
            return next('/login');
        }
    }

    if (to.matched.some(record => record.meta.guestOnly)) {
        // Login/Register pages
        if (token) {
            return next('/');
        }
    }
    next();
});

export default router;
