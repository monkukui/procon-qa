<template>
  <div class="home">
    <TopInfo />
    <div class="left">
      <LeftInfo />
    </div>
    <Questions />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Questions from '@/components/Questions.vue';
import TopInfo from '@/components/TopInfo.vue';
import LeftInfo from '@/components/LeftInfo.vue';

@Component({
  components: {
    TopInfo,
    LeftInfo,
    Questions,
  },
})
export default class Home extends Vue {
}
</script>

<style scoped>
.left {
  postion: -webkit-sticky;
  position: sticky;
  top: 20%;
  left: 20%;
}

</style>
