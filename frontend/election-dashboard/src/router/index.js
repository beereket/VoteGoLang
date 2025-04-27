import Vue from 'vue';
import Router from 'vue-router';
// import { jwtDecode } from 'jwt-decode';
import LoginPage from "@/views/auth/LoginPage.vue";
import CandidatesPage from "@/views/candidates/Candidates.vue";
import DashboardAnalytics from "@/views/DashboardAnalytics.vue";
import PetitionsPage from "@/views/petitions/PetitionsPage.vue";
import PetitionVotingAnalytics from "@/views/petitions/PetitionVotingAnalytics.vue";
import DashboardPage from "@/views/Dashboard.vue";
import PetitionCommentsPage from "@/views/petitions/PetitionCommentsPage.vue";
import NewsManagement from "@/views/admin/NewsManagement.vue";
import UserManagement from "@/views/admin/UserManagement.vue";
import ActivityLogs from "@/views/admin/ActivityLogs.vue";


Vue.use(Router);

const router = new Router({
    mode: 'history',
    routes: [
        { path: '/login', component: LoginPage },
        {
            path: '/admin/dashboard',
            component: DashboardPage,
            meta: { requiresAdmin: true },
        },
        {
            path: '/admin/candidates',
            component: CandidatesPage,
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
            path: '/admin/petitions',
            component: PetitionsPage,
            meta: { requiresAdmin: true },
        },
        {
            path: '/admin/petitions/comments',
            component: PetitionCommentsPage,
            meta: { requiresAdmin: true },
        },
        {
            path: '/admin/petitions/voting-analytics',
            component: PetitionVotingAnalytics,
            meta: { requiresAdmin: true },
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
        if (!token || role !== 'admin') {
            return next('/login');
        }
    }

    next();
});

export default router;
