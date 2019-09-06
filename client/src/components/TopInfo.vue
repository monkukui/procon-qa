<template>
  <div class="top">
    <v-btn
      depressed large 
      color="primary"
      @click="openModal"
    >質問する</v-btn>
    <QuestionModal
      v-if="canDisplayModal"
      @closeModal="closeModal"
    />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit, Watch} from 'vue-property-decorator';
import OriginalButton from '@/components/OriginalButton.vue';
import QuestionModal from '@/components/QuestionModal.vue';

@Component({
  components: {
    OriginalButton,
    QuestionModal,
  },
})
export default class TopInfo extends Vue {

  private canDisplayModal: boolean = false;
  private openModal(): void {
    this.canDisplayModal = true;
  }
  private closeModal(): void {
    this.canDisplayModal = false;
  }
}
</script>

<style scoped>
</style>
