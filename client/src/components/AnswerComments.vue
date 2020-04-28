<template>
  <div>
    <div v-for="(comment, index) in comments" :key=index>
      <CommentPanel
        :uid="comment.uid"
        :body="comment.body"
        :date="comment.date"
      />
    </div>

    <CommentForm
      @comment="postComment"
      :qid="qid"
      :aid="aid"
      :uid="uid"
      type="3"
    />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';

import CommentPanel from '@/components/CommentPanel.vue';
import CommentForm from '@/components/CommentForm.vue';

@Component({
  components: {
    CommentPanel,
    CommentForm,
  },
})
export default class AnswerComments extends Vue {
  @Prop()
  private aid!: string;
  @Prop()
  private uid!: string;
  @Prop()
  private type!: string;
  @Prop()
  private qid!: string;

  private comments: any[] = [];

  private postComment(): void {
    this.getComments();
  }

  private getComments(): void {
    const url = '/api/no-auth/answer-comment/' + this.aid;
    fetch(url).then((response) => {
      if (response.ok) {
        return response.json();
      } else {
        alert('error');
      }
      return [];
    }).then((json) => {
      this.comments = [];
      this.comments = json;
    });
  }

  private created() {
    this.getComments();
  }
}
</script>
