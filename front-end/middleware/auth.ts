export default defineNuxtRouteMiddleware((to, from) => {
    const authStore = useAuth();


    if (!authStore.user) {
        return navigateTo('/login');
    }
});