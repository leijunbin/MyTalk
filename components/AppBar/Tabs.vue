<template>
  <v-tabs
    class="Tabs"
    slider-color="#22469a"
    :centered="true"
    :background-color="$vuetify.theme.dark ? '#121212' : ''"
  >
    <v-tab v-for="route in routes" :key="route.key" :to="route.to" nuxt>
      {{ route.title }}
    </v-tab>
  </v-tabs>
</template>

<script>
import routes from "../../config/routes";
export default {
  data() {
    return {
      routes,
    };
  },
};
</script>

<style lang="scss">
.Tabs {
  height: 100%;
  .v-tab {
    font-weight: bold;
  }
}
</style>