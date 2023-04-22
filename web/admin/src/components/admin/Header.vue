<template>
    <a-button type="danger" @click="loginout">退出</a-button>
</template>

<script>
export default {
    methods: {
        loginout() {
            window.sessionStorage.clear('token')
            this.$router.push('/login')
        }
    }
}
</script>

<style scoped></style>