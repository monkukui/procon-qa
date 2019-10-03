<template>
  <div class="about">
    <h1>UI テストページ</h1>
    <QuestionPanel />
    <QuestionPanelDetail />
    <AnswerPanelDetail />
    <AnswerForm />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import QuestionPanelDetail from '@/components/QuestionPanelDetail.vue';
import QuestionPanel from '@/components/QuestionPanel.vue';
import AnswerPanelDetail from '@/components/AnswerPanelDetail.vue';
import AnswerForm from '@/components/AnswerForm.vue';

@Component({
  components: {
    QuestionPanel,
    QuestionPanelDetail,
    AnswerPanelDetail,
    AnswerForm,
  },
})
export default class About extends Vue {
}
</script>

<style>
</style>
