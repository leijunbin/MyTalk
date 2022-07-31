<template>
  <!-- v-dialog:对话框 persistent:点击元素外面或按下 esc 键将不会关闭它-->
  <!-- fullscreen:对话框是否铺满屏幕，网站策略是对话框在手机设备时全屏显示-->
  <!-- vuetify.breakpoint:显示大小监听钩子，name为现在大小等级，xs为手机大小等级-->
  <!-- transition:对话框的弹出和退出动画样式，dialog-bottom-transition为从底部弹出-->
  <v-dialog
    max-width="490"
    v-model="AccountDialog"
    persistent
    :fullscreen="$vuetify.breakpoint.name === 'xs'"
    transition="dialog-bottom-transition"
  >
    <v-card>
      <div class="back-header">
        <v-btn elevation="0" text @click="SET_ACCOUNT_DIALOG">
          <v-icon>mdi-arrow-left</v-icon>
        </v-btn>
      </div>
      <div class="account-box">
        <header>
          <h2>
            {{ topTitle }}
            <v-btn
              text
              color="red accent-2"
              class="header_action-btn"
              @click="type = !type"
            >
              {{ typeToName }}
            </v-btn>
          </h2>
          <h3>请填写以下信息进行{{ retypeToName }}</h3>
        </header>
        <AccountLogin v-show="type" />
        <AccountRegister v-show="!type" />
      </div>
    </v-card>
  </v-dialog>
</template>

<script>
import { mapMutations, mapState } from "vuex";
import AccountLogin from "./AccountLogin.vue";
import AccountRegister from "./AccountRegister.vue";
export default {
  components: {
    AccountLogin,
    AccountRegister,
  },
  data() {
    return {
      type: true,
    };
  },
  computed: {
    ...mapState(["AccountDialog"]),
    topTitle() {
      return this.type ? "欢迎回来，" : "欢迎您，";
    },
    typeToName() {
      return this.type ? "注册" : "登录";
    },
    retypeToName() {
      return this.type ? "登录" : "注册";
    },
  },
  methods: {
    ...mapMutations(["SET_ACCOUNT_DIALOG"]),
  },
};
</script>

<style lang="scss">
.back-header {
  padding: 8px 6px 0; // 间距，此种写法的顺序使上、左右、下
}
.account-box {
  padding: 12px 28px 34px;
  h2 {
    display: flex;
    align-content: center;
    justify-content: space-between;
    margin-bottom: 6px;
  }
  h3 {
    color: #999;
    margin-bottom: 12px;
  }
}
</style>