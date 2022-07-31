<template>
  <v-form class="main-form">
    <div class="form-field input-required">
      <label>邮箱</label>
      <v-text-field v-model="email" filled rounded dense required />
    </div>
    <div class="form-field input-required">
      <label>密码</label>
      <v-text-field
        :type="hidePassword ? 'text' : 'password'"
        filled
        rounded
        dense
        required
        v-model="password"
      >
        <template slot="append-outer">
          <v-btn small rounded text @click="hidePassword = !hidePassword">
            <v-icon size="20px">{{ hidePasswordIcon }}</v-icon>
          </v-btn>
        </template>
      </v-text-field>
    </div>
    <v-btn
      large
      class="next-btn"
      elevation="0"
      rounded
      :loading="loading"
      :color="timer ? 'error' : ''"
    >
      <span class="subtitle-2" v-if="timer"
        >操作频繁，请等待 {{ second }} 秒</span
      >
      <span class="subtitle-2" v-show="timer === null">继续</span>
      <v-icon size="18px" class="ml-2" v-show="timer === null"
        >mdi-arrow-right</v-icon
      >
    </v-btn>
  </v-form>
</template>

<script>
export default {
  data() {
    return {
      hidePassword: true,
      email: "",
      password: "",
      loading: false,
      timer: null,
    };
  },
  computed: {
    hidePasswordIcon() {
      return this.hidePassword ? "mdi-eye-off" : "mdi-eye";
    },
  },
};
</script>

<style lang="scss">
// 暗色模式样式调整
.theme--dark {
  .main-form {
    .form-field {
      label {
        color: #9e9e9e !important;
      }
      .v-text-field input {
        background-color: #272727;
        color: #fff;
      }
      .v-text-field input:hover {
        background-color: #121212;
      }
    }
  }
}
.main-form {
  margin-top: 30px;
  .next-btn {
    width: 100%;
    margin: 16px auto 0;
  }
  .form-field {
    .v-text-field input {
      width: 100%;
      font-size: 14px; // 输入框字体大小
      padding: 10px 16px; // 输入内容距离输入框距离
      border-radius: 8px; // 输入框圆角大小
      background-color: #f3f3f4;
      max-height: inherit; // 段落最大高度，inherit从父元素继承max-height的值
    }
    .v-text-field input:hover {
      background-color: #fff;
      border-color: rgba(4, 120, 190, 0.4);
      box-shadow: 0 0 0 4px rgba(4, 120, 190, 0.1);
    }
    .v-input__append-outer {
      margin-top: 6px !important; // !important 覆盖父组件相同参数
    }
    .v-input__slot {
      padding: 0;
    }
  }
  .form-field.input-required {
    label {
      padding-left: 8px;
      display: block;
      font-weight: 700;
      margin: 14px 0 6px;
      color: #0d0c22;
      &::after {
        content: "*";
        color: red;
        margin-left: 4px;
      }
    }
  }
}
</style>