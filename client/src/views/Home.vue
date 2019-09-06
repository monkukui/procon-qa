<template>
  <div class="home">
    <div class="main-contents">
      <Questions />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Questions from '@/components/Questions.vue';
import TopInfo from '@/components/TopInfo.vue';

@Component({
  components: {
    TopInfo,
    Questions,
  },
})
export default class Home extends Vue {
}
</script>

<style scoped>
.home {
  margin: 5%;
}
</style>
