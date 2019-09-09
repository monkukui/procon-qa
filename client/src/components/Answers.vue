<template>
  <div class="answers">
    
    <AnswerPanelDetail
      body="それは自明では?"
      answeredTime="2019/05.26 11:11:21"
    />
    <AnswerPanelDetail
      body="自明でしょ"
      answeredTime="2019/05.26 11:11:21"
    />
    <AnswerPanelDetail 
      body="流石に自明すぎわロタ"
      answeredTime="2019/05.26 11:11:21"
    />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit, Watch} from 'vue-property-decorator';
import AnswerPanelDetail from '@/components/AnswerPanelDetail.vue';
@Component({
  components: {
    AnswerPanelDetail,
  },
})
export default class Answers extends Vue {
  private numOfAnswers: number = 3;
}
</script>

<style scoped>
.answers {
  margin: 0px 0px 20px-24px;
}
</style>
