<template>
  <div>
    <v-divider></v-divider>
    <v-row>
      <v-col cols="12" sm="8">
        <p>{{ body }}</p>
      </v-col>
      <v-col cols="12" sm="4">
        <v-row>
          <v-col cols="12" sm="4">
            <TwitterIcon
              :twitterId="twitterId"
              :uid="uid"
              size="36"
              apiSize="n"
            />
          </v-col>
          <v-col cols="12" sm="8">
            <UserName
              :name="name"
              :uid="uid"
            />
          </v-col>
        </v-row>
        <p>{{ date }}</p>
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

  private name: string = '';
  private twitterId: string = '';

  private created(): void {
    this.setUser();
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
}
</script>
