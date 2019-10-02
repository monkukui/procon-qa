<template>
  <div class="about">
    <h1>UI テストページ</h1>
    <QuestionPanel />
    <QuestionPanelDetail />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import QuestionPanelDetail from '@/components/QuestionPanelDetail.vue';
import QuestionPanel from '@/components/QuestionPanel.vue';

@Component({
  components: {
    QuestionPanel,
    QuestionPanelDetail,
  },
})
export default class About extends Vue {
}
</script>

<style>
</style>
