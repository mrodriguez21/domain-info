<template>
  <b-form-group
    id="fieldset"
    :invalid-feedback="invalidFeedback"
    :valid-feedback="validFeedback"
    :state="state"
  >
    <b-input-group prepend="Domain: ">
      <b-form-input
        id="domainName"
        v-model="domainName"
        @keyup.enter="onSearchSubmition"
        :state="state"
        :maxlength="50"
        trim
      ></b-form-input>
      <b-input-group-append>
        <span class="input-group-text">
          <i class="fa fa-search fa-lg"></i>
        </span>
      </b-input-group-append>
    </b-input-group>
  </b-form-group>
</template>

<script>
export default {
  computed: {
    state() {
      if (this.domainName.length == 0) return null;
      // Regular expression taken from: https://www.socketloop.com/tutorials/golang-use-regular-expression-to-validate-domain-name
      let regExp = new RegExp(
        /^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$/
      );
      return regExp.test(this.domainName);
    },
    invalidFeedback() {
      return this.state === true ? "" : "Please enter a valid domain name.";
    },
    validFeedback() {
      return this.state === true
        ? "This is a valid domain. Now press Enter."
        : "";
    },
  },
  data() {
    return {
      domainName: "",
    };
  },
  methods: {
    onSearchSubmition() {
      if (this.state) {
        this.$emit("search-submitted", this.domainName.toLowerCase());
      }
    },
  },
};
</script>

<style>
#fieldset {
  width: 60vw;
  margin: auto;
  margin-bottom: 40px;
}
</style>
