<template>
  <a-layout class="container">
    <Nav></Nav>
    <a-layout>
      <a-layout-header class="headerbtn">
        <Header></Header>
      </a-layout-header>
      <a-layout-content>
        <router-view :key="$route.path"></router-view>
      </a-layout-content>
      <a-layout-footer>
        <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import Nav from '../components/admin/Nav.vue'
import Footer from '../components/admin/Footer.vue'
import Header from '../components/admin/Header.vue'
export default {
  name: "Admin",
  components: { Nav, Footer, Header },
}
</script>

<style scoped>
.container {
  height: 100%;
}

.headerbtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>