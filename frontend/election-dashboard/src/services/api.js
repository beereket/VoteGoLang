import axios from 'axios';

// 1️⃣ Create Axios instance
const api = axios.create({
    baseURL: 'http://localhost:8080',
    headers: {
        'Content-Type': 'application/json',
    },
});

api.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token');
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`;

            console.log('✅ Authorization header attached:', config.headers['Authorization']);
        } else {
            console.log('❌ No token found in localStorage');
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

export default api;
