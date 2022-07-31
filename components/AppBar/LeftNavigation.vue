<template>
  <v-navigation-drawer app temporary v-model="leftNavigationComputed">
    <div class="nav-bar-header">
      <div class="nav-bar-header-btn-group">
        <v-btn depressed icon color="#f1e8eb" @click="handleChangeTheme">
          <v-icon>{{ darkModeIcon }}</v-icon>
        </v-btn>
        <v-btn depressed icon color="#f1e8eb">
          <v-icon>mdi-bell</v-icon>
        </v-btn>
      </div>
      <div class="nav-bar-header-avatar">
        <v-avatar size="100">
          <v-img :src="require('~/static/icon.png')" alt="" />
        </v-avatar>
        <v-list-item-title class="title">
          <v-icon>mdi-account</v-icon>
          <span class="username">{{ username }}</span>
        </v-list-item-title>
      </div>
    </div>
    <div class="nav-container">
      <div class="nav-content">
        <v-list>
          <v-list-item-group>
            <v-list-item
              v-for="route in routes"
              :to="route.to"
              :key="route.key"
              class="rounded"
              nuxt
            >
              <v-list-item-icon>
                <v-icon>{{ route.icon }}</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>{{ route.title }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </div>
    </div>
    <div slot="append" class="container nav-footer">
      <v-btn block rounded color="blue" @click="SET_ACCOUNT_DIALOG">登录</v-btn>
      <!-- <div class="authority" v-else>
          <v-btn text>
            <v-icon>mdi-cog</v-icon>
            设置
          </v-btn>
          <v-btn text @click="onLogout">
            <v-icon>mdi-exit-to-app</v-icon>
            退出
          </v-btn>
        </div> -->
    </div>
  </v-navigation-drawer>
</template>

<script>
import { mapMutations, mapState } from "vuex";
import routes from "../../config/routes";
export default {
  data() {
    return {
      routes,
      username: "barar",
    };
  },
  computed: {
    ...mapState(["LeftNavigation"]),
    leftNavigationComputed: {
      get() {
        return this.LeftNavigation;
      },
      set(val) {
        this.$store.commit("SET_SIDE_STATUS", val);
      },
    },
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
  },
};
</script>

<style lang="scss">
.nav-bar-header {
  width: 100%;
  display: flex;
  flex-direction: column;
  &-btn-group {
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    padding: 12px;
  }
  &-avatar {
    padding: 24px 0;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    .title {
      margin-top: 18px;
      font-weight: bold;
      display: flex;
      flex-direction: row;
      .username {
        display: inline-block;
        margin-left: 4px;
      }
    }
  }
}

.nav-container {
  .nav-content {
    border-radius: 30px 30px 0 0;
    padding: 30px 0;
    .v-list {
      margin: 0px auto;
      width: 200px;
    }
    .v-list-item {
      margin-bottom: 8px;
      min-height: 40px;
      border-radius: 32px !important;
      &__icon {
        margin: 8px 0;
        margin-right: 32px;
      }
      &__content {
        padding: 4px 0;
      }
    }
    .v-list-item:focus::before,
    .v-list-item:hover::before,
    .v-list-item--active::before {
      opacity: 0 !important;
      .v-list-item-icon.v-icon {
        color: #ffffff;
      }
    }
  }
}

.authority {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  .v-btn .v-icon {
    margin-right: 8px;
  }
  .v-btn.theme--light {
    color: #ffffff;
  }
  .v-btn.theme--dark {
    color: #ffffff;
  }
}

.v-list-item__title {
  color: #ffffff;
}
.v-navigation-drawer__content {
  a.v-list-item,
  .v-icon {
    color: #ffffff;
  }
}
.container.nav-footer .v-btn__content {
  color: #fff;
}

.theme--dark {
  .nav-bar-header {
    background-color: #0b0b0b;
  }
  .nav-container {
    background-color: #0b0b0b;
    .nav-content {
      background-color: #1d1d1e;
      .v-list.v-sheet {
        background-color: #1d1d1e;
      }
    }
  }
  .v-navigation-drawer__content {
    background-color: #1d1d1e;
  }
  .container.nav-footer {
    background-color: #1d1d1e;
  }
}

.theme--light {
  .nav-bar-header {
    background-color: #242663;
  }
  .nav-container {
    background-color: #242663;
    .nav-content {
      background-color: #1a1d53;
      .v-list.v-sheet {
        background-color: #1a1d53;
      }
    }
  }
  .v-navigation-drawer__content {
    background-color: #1a1d53;
  }
  .container.nav-footer {
    background-color: #1a1d53;
  }
}
</style>