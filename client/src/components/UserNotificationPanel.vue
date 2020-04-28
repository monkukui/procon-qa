<template>
  <span>
    <v-divider
      v-if="shouldDivide"
    />
    <v-list three-line>
      <template>
        <v-list-item
          :to="{ name: 'question', query: { questionId: this.qid }}"
        >
          <span class="icon">
            <TwitterIcon
              size="48"
              apiSize="b"
              :twitterId="twitterId"
              :uid="uid"
            />
            <UserName
              :uid="uid"
              :name="name"
            />
          </span>
          <v-list-item-content>
            <v-list-item-title v-html="title"></v-list-item-title>
            <v-list-item-subtitle v-html="body"></v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </template>
    </v-list>
  </span>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit, Watch } from 'vue-property-decorator';
import UserName from '@/components/UserName.vue';
import TwitterIcon from '@/components/TwitterIcon.vue';

@Component({
  components: {
    UserName,
    TwitterIcon,
  },
})
export default class QuestionPanel extends Vue {
  @Prop({required: true})
  private qid!: string;
  @Prop({required: true})
  private ouid!: string;
  @Prop({required: true})
  private type!: number;
  @Prop({required: true})
  private body!: string;
  @Prop({required: true})
  private shouldDivide!: boolean;

  private title: string = '';

  private name: string = '';
  private twitterId: string = '';

  private created(): void {
    this.setUser();

    if (this.type === 1) {
      this.title = 'あなたの質問に回答がつきました';
    } else if (this.type === 2) {
      this.title = 'あなたの質問にコメントがつきました';
    } else if (this.type === 3) {
      this.title = 'あなたの回答にコメントがつきました';
    }
  }

  private setUser(): void {

    const url = 'api/no-auth/user/' + this.ouid;
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.name = json.name;
      this.twitterId = json.twitter_id;
    });
  }

  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>

<style scoped>
.icon {
  margin-right: 30px;
}
</style>
