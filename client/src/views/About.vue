<template>
  <div class="about">
    <TopInfo />
    <h1>This is an about page</h1>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import TopInfo from '@/components/TopInfo.vue';

@Component({
  components: {
    TopInfo,
  },
})
export default class About extends Vue {
}
</script>

