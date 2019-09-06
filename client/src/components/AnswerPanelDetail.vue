<template>
  <div class="answer-panel-detail">
    <div class="box">
      <h4>回答</h4>
      <hr>
      <p>{{ body }}</p>
      <span class="small">回答日時: {{ answeredTime }}</span>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';

@Component
export default class AnswerPanelDetail extends Vue {
  @Prop()
  private body!: string;
  @Prop()
  private answeredTime!: string;
}
</script>

<style scoped>

.small {
  font-size: 75%;
}

.answer-panel-detail {
  font-size: 75%;
}

.box{
  width: 75%;
  padding: 0.5em 1em;
  margin: 2em 0;
  color: #5d627b;
  background: white;
  border-top: solid 10px #5d627b;
  box-shadow: 0 3px 5px rgba(0, 0, 0, 0.22);
}
.box p {
  margin: 0; 
  padding: 0;
}
  
</style>
