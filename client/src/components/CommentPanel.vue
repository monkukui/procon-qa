<template>
  <div>
    <v-divider></v-divider>
    <v-row>
      <v-col cols="12" sm="9">
        <p>{{ body }}</p>
        <span class="grey--text lighten-4">{{ date }}</span>
      </v-col>
      <v-col cols="12" sm="3">
        <div class="text-right">
          <TwitterIcon
            :twitterId="twitterId"
            :uid="uid"
            size="36"
            apiSize="n"
          />
          <UserName
            :name="name"
            :uid="uid"
          />
          <br>
          <v-icon v-if="name===userName" color="error" @click="deleteComment">mdi-delete</v-icon>
        </div>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import UserName from '@/components/UserName.vue';
import TwitterIcon from '@/components/TwitterIcon.vue';

@Component({
  components: {
    UserName,
    TwitterIcon,
  },
})
export default class CommentPanel extends Vue {
  @Prop()
  private uid!: string;
  @Prop()
  private body!: string;
  @Prop()
  private date!: string;
  @Prop()
  private id!: string;
  @Prop()
  private type!: string; // "2" -> for question, "3" -> for answer

  private name: string = '';
  private userName: string = '';
  private twitterId: string = '';

  private created(): void {
    this.setUser();
    if (this.getToken() != null) {
      const claims = JSON.parse(atob(this.getToken().split('.')[1]));
      this.userName = claims.name;
    }
  }

  private deleteComment(): void {

    const url = 'api/' + (this.type === '2' ? 'question' : 'answer') + '-comment/' + this.id;
    const method = 'DELETE';
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {method, headers}).then((response) => {
      if (response.ok) {
        // 質問を削除しました
        this.$emit('comment');
      }
    });
  }

  private setUser(): void {
    const url = 'api/no-auth/user/' + this.uid;
    fetch(url).then((response) => {
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
