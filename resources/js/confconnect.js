import axios from 'axios';

const instance = axios.create({
    //baseURL: '/', 
    baseURL: 'http://127.0.0.1:8093/',
    headers: {
        'Content-type': 'application/json',
    },

});

instance.interceptors.request.use(config => {
    const token = localStorage.getItem('token'); 
    if (token) {
        config.headers['Authorization'] = `Bearer ${token}`; 
    }
    return config;
}, error => {
    return Promise.reject(error);
});

export default instance;

