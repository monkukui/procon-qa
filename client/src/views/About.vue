<template>
  <div class="about">
    <h1>UI テストページ</h1>
    <QuestionPanel />
    <QuestionPanelDetail />
    <AnswerPanelDetail />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import QuestionPanelDetail from '@/components/QuestionPanelDetail.vue';
import QuestionPanel from '@/components/QuestionPanel.vue';
import AnswerPanelDetail from '@/components/AnswerPanelDetail.vue';

@Component({
  components: {
    QuestionPanel,
    QuestionPanelDetail,
    AnswerPanelDetail,
  },
})
export default class About extends Vue {
}
</script>

<style>
</style>
