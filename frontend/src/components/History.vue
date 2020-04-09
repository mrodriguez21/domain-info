<template>
  <div class="container">
    <ul>
      <Domain
        v-for="item in items"
        :key="item.domain"
        :name="item.domain"
        :domain="item.info"
      ></Domain>
    </ul>
  </div>
</template>

<script>
import Domain from "@/components/Domain.vue";
import APIClient from "../apiClient";

export default {
  name: "History",
  components: {
    Domain,
  },
  data() {
    return {
      items: [],
    };
  },
  created() {
    this.getHistory();
  },
  methods: {
    getHistory() {
      APIClient.getPreviousDomains().then((response) => {
        this.items = response.items;
      });
    },
  },
};
</script>
