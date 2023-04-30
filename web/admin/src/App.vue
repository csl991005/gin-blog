<template>
  <!-- <a-locale-provider :locale="locale"> -->
  <a-config-provider :locale="locale">
    <div id="app">
      <router-view></router-view>
    </div>
  </a-config-provider>
  <!-- </a-locale-provider> -->
</template>

<script>
import zhCN from 'ant-design-vue/es/locale/zh_CN';
export default {
  data() {
    return {
      locale: zhCN,
    };
  },
};
</script>