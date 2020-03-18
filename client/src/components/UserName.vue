<template>
  <span v-bind:class="color">
    <span v-if="name">
      {{ name }}
    </span>
    <span v-else>
      退会済みのユーザ
    </span>
  </span>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class UserName extends Vue {
  @Prop()
  private name!: string;
  @Prop()
  private color!: string;
}
</script>

<style>
.user-red {color:#FF0000;}
.user-orange {color:#FF8000;}
.user-yellow {color:#C0C000;}
.user-blue {color:#0000FF;}
.user-cyan {color:#00C0C0;}
.user-green {color:#008000;}
.user-brown {color:#804000;}
.user-gray {color:#808080;}
.user-unrated {color:#000000;}
.user-admin {color:#C000C0;}
.username>span {font-weight:bold;}
</style>
