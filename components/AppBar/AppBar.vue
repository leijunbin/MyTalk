<template>
  <div class="AppBar">
    <LeftNavigation />
    <!--v-app-bar:应用栏 app:指定该组件作为应用程序一部分 dense:将应用栏高度降低到48px flat:删除应用栏阴影 height:指定工具栏特定高度,包括dense tile:删除组件圆角样式 -->
    <v-app-bar app flat height="56px" tile fixed>
      <div class="toolbar-content">
        <div class="toolbar-title">
          <!-- v-btn:按钮 elevation:组件在z轴上的深度(影响组件阴影) text:使背景变为透明 href:指定组件为锚点并应用href属性 nuxt:指定链接nuxt-link -->
          <v-btn elevation="0" text href="/" nuxt>MyTalk</v-btn>
        </div>
        <Tabs></Tabs>
        <!-- small:按钮尺寸变小 -->
        <v-btn
          class="themeBtn"
          elevation="0"
          text
          small
          depressed
          @click="handleChangeTheme"
        >
          <v-icon>{{ darkModeIcon }}</v-icon>
        </v-btn>
        <v-btn
          class="loginBtn"
          color="info"
          element="0"
          depressed
          @click="SET_ACCOUNT_DIALOG"
          >登录</v-btn
        >
      </div>
      <v-app-bar-nav-icon
        class="leftNavigationIcon"
        @click.stop="changeLeftNavigation"
      ></v-app-bar-nav-icon>
    </v-app-bar>
  </div>
</template>

<script>
import { mapMutations } from "vuex";
import Tabs from "../AppBar/Tabs.vue";
import LeftNavigation from "../AppBar/LeftNavigation.vue";

export default {
  components: { Tabs, LeftNavigation },
  data() {
    return {
      title: "MyTalk",
    };
  },
  computed: {
    darkModeIcon() {
      return this.$vuetify.theme.dark
        ? "mdi-white-balance-sunny"
        : "mdi-weather-night";
    },
  },
  methods: {
    ...mapMutations(["SET_ACCOUNT_DIALOG"]),
    handleChangeTheme() {
      this.$vuetify.theme.dark = !this.$vuetify.theme.dark;
    },
    changeLeftNavigation() {
      this.$store.commit("SET_SIDE_STATUS", true);
    },
  },
};
</script>

<style lang="scss">
.theme--light.v-application {
  background-color: #f4f4f4;
}
.theme--light.v-app-bar.v-toolbar.v-sheet {
  background-color: #fff;
}
.theme--dark.v-app-bar.v-toolbar.v-sheet {
  background-color: #121212;
}
.AppBar {
  @media (max-width: 824px) {
    .toolbar-content {
      .themeBtn {
        display: none !important;
      }
      .loginBtn {
        display: none !important;
      }
      .Tabs {
        display: none !important;
      }
    }
    .leftNavigationIcon {
      display: block !important;
    }
  }
  .leftNavigationIcon {
    display: none;
  }
  .toolbar-content {
    width: 100%;
    height: 100%;
    margin: 0 auto; // 设置对象外边距，两个参数 top和bottom left和right auto:自适应相同值，即居中
    display: flex; // 弹性布局，可以拓展和伸缩容器内的元素，最大限度利用容器空间
    align-items: center; // 对齐弹性盒各项方式：center 居中对齐
    justify-content: space-between; // 弹性盒中元素的空白方式，space-between 就是均匀排列每个元素，首个元素位于起点，末尾元素位于终点
    .toolbar-title {
      .v-btn {
        font-size: 20px;
        font-weight: bold;
      }
    }
    .v-btn {
      margin-right: 12px;
    }
  }
}
</style>