<template>
  <div class="container">
    <div class="datainfo">
      <div class="status">
        <mdb-tooltip trigger="hover" :options="{ placement: 'left' }">
          <span slot="tip">
            Current SSL grade
          </span>
          <div slot="reference" class="ssl_grade mb-1">
            <i class="ssl_icon fas fa-circle fa-2x" :class="iconColor"></i>
            <span :class="textColor">{{ domain.ssl_grade || "-" }}</span>
          </div>
        </mdb-tooltip>
        <mdb-tooltip trigger="hover" :options="{ placement: 'left' }">
          <span slot="tip">
            Previous SSL grade
          </span>
          <div slot="reference" class="ssl_grade">
            <i class="ssl_icon fas fa-circle fa-2x" :class="prevIconColor"></i>
            <span :class="prevTextColor">
              {{ domain.previous_ssl_grade || "-" }}
            </span>
          </div>
        </mdb-tooltip>
      </div>
      <div class="fav">
        <img
          v-if="domain.logo != ''"
          class="domain-logo"
          :src="domain.logo"
          alt="Logo"
        />
        <img
          v-else
          class="domain-logo"
          src="@/assets/images/question-circle-solid.png"
          alt="No logo found"
        />
      </div>
      <div class="info">
        <div class="subinfo">
          <span class="domain-name title">{{ name }}</span>
          <div>
            <mdb-tooltip trigger="hover" :options="{ placement: 'right' }">
              <span slot="tip">
                Host is currently
                <b>{{ domain.is_down ? "down" : "up" }}.</b>
              </span>
              <i
                slot="reference"
                class="fas ml-2 fa-2x"
                :class="[
                  domain.is_down
                    ? 'fa-times-circle red-text'
                    : 'fa-check-circle green-text',
                ]"
              ></i>
            </mdb-tooltip>
          </div>
        </div>
        <div>{{ domain.title }}</div>
      </div>
      <div
        class="toggle"
        v-if="!domain.is_down && collapsed"
        v-b-toggle="'collapse-' + name"
      >
        <i class="fas fa-chevron-down"></i>
      </div>
    </div>
    <ServersTable
      v-if="!domain.is_down && !collapsed"
      :servers="domain.servers"
      :servers_changed="domain.servers_changed"
    ></ServersTable>
    <b-collapse
      v-else-if="!domain.is_down && collapsed"
      :id="'collapse-' + name"
      class="mt-2"
    >
      <ServersTable
        :servers="domain.servers"
        :servers_changed="domain.servers_changed"
      ></ServersTable>
    </b-collapse>
  </div>
</template>

<script>
import { mdbTooltip } from "mdbvue";
import ServersTable from "@/components/ServersTable.vue";
export default {
  components: { mdbTooltip, ServersTable },
  computed: {
    iconColor() {
      switch (this.domain.ssl_grade) {
        case "A+":
          return "green-text";
        case "A":
          return "light-green-text";
        case "A-":
          return "lime-text";
        case "B":
          return "yellow-text";
        case "C":
          return "amber-text";
        case "D":
          return "orange-text";
        case "E":
          return "deep-orange-text";
        case "F":
          return "red-text";
        default:
          return "text-light";
      }
    },
    prevIconColor() {
      switch (this.domain.previous_ssl_grade) {
        case "A+":
          return "green-text";
        case "A":
          return "light-green-text";
        case "A-":
          return "lime-text";
        case "B":
          return "yellow-text";
        case "C":
          return "amber-text";
        case "D":
          return "orange-text";
        case "E":
          return "deep-orange-text";
        case "F":
          return "red-text";
        default:
          return "text-light";
      }
    },
    textColor() {
      return this.domain.ssl_grade == "B" ? "text-black-50" : "white-text";
    },
    prevTextColor() {
      return this.domain.previous_ssl_grade == "B"
        ? "text-black-50"
        : "white-text";
    },
  },
  props: {
    name: String,
    domain: Object,
    collapsed: Boolean,
  },
};
</script>

<style>
.datainfo {
  display: flex;
  flex: 1;
  flex-direction: row;
  justify-items: center;
  align-items: stretch;
  width: 60vw;
  margin: auto;
  margin-bottom: 16px;
  border: 1px solid lightgray;
  padding: 20px;
  border-radius: 10px;
  /*! box-shadow: 0 1px 13px -7px black; */
}
.status {
  flex-direction: column;
  display: flex;
  justify-content: space-around;
  margin-right: 10px;
}
.fav {
  align-items: center;
  display: flex;
}
.info {
  display: flex;
  flex-direction: column;
  flex: 5;
  flex-grow: 1;
}
.subinfo {
  flex-direction: row;
  display: flex;
  flex-grow: 1;
  align-items: center;
  justify-items: self-end;
  align-self: center;
}
.subinfo > div {
  justify-self: center;
  align-self: center;
}
.ssl_grade {
  position: relative;
  display: flex;
  height: 2em;
  width: 2em;
}
.ssl_grade > i {
  position: absolute;
}
.ssl_grade > span {
  margin: auto;
  z-index: 10;
  font-weight: bold;
}

.domain-logo {
  height: 4.25rem;
}
.domain-name {
  padding-bottom: 5px;
}

.toggle {
  display: flex;

  align-items: center;

  cursor: pointer;
}
.toggle > i {
  transition: transform 0.4s ease-in-out;
}
.toggle:not(.collapsed) > i {
  transform: rotate(180deg);
}

.title {
  font-size: 1.5rem;
}
</style>
