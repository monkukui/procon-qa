<template>
  <div>
    <h3>回答する</h3>
    <textarea></textarea>
  </div>
</template>

<script lang="ts">
import {Component, Emit, Prop, Vue} from 'vue-property-decorator';

@Component({
  components: {
  },
})

export default class QuestionModal extends Vue {
}
</script>

<style>

textarea {
  font-size: 18px;
  width: 100%;
  height: 180px;
  resize: none;
}

</style>
