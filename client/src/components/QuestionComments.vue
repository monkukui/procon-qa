<template>
  <div>
    <span v-for="(comment, index) in comments" :key=index>
      <CommentPanel
        :uid="comment.uid"
        :body="comment.body"
        :date="comment.date"
      />
    </span>
    <v-row>
      <v-col
        cols="auto"
        class="mr-auto"
      >
      </v-col>
      <v-col cols="auto">
        <v-btn
          class="ma-2"
          color="indigo"
          text
          @click="commentFormOpen = !commentFormOpen"
        >
          コメントを追加する
        </v-btn>
      </v-col>
    </v-row>
    <CommentForm 
      v-if="commentFormOpen"
      @comment="postComment"
      :qid="qid"
      aid="-1"
      :uid="uid"
      type="2"
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
export default class QuestionComments extends Vue {
  @Prop()
  private qid!: string;
  @Prop()
  private uid!: string;

  private xor: boolean = false;
  private commentFormOpen: boolean = false;

  private comments: any[] = [];

  private postComment(): void {
    this.getComments();
  }
  private getComments(): void {
    const url = '/api/no-auth/question-comment/' + this.qid;
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
