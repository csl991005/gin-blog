<template>
    <v-app app>
        <TopBar></TopBar>

        <v-main>
            <v-container></v-container>
        </v-main>

        <Footer></Footer>
    </v-app>
</template>

<script>
import TopBar from '@/components/TopBar.vue'
import Footer from '@/components/Footer.vue'
export default {
    name: 'Home',

    components: { TopBar, Footer }
}
</script>
