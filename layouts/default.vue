<template>
  <v-app>
    <AppBar />
    <!-- 页面主要布局 -->
    <v-main>
      <!-- 提供了网站内容居中和水平填充功能，可以使用 fluid 属性将容器在所有视口和设备尺寸上完全拓展。 -->
      <v-container fluid>
        <!-- 引入nuxt路由 -->
        <nuxt />
      </v-container>
    </v-main>
    <!-- 页脚 -->
    <v-footer app absolute>123</v-footer>
    <Account />
  </v-app>
</template>

<script>
import Account from "@/components/Account/Account.vue";
import AppBar from "@/components/AppBar/AppBar.vue";

export default {
  components: {
    Account,
    AppBar,
  },
};
</script>

<style lang="scss">
</style>