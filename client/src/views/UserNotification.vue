<template>
  <div v-if="uid">
    <h1>通知</h1>
    <div v-if="isReady">
      {{ notifiactions }}
      <v-list two-line>
        <template v-for="(item, index) in items.slice(0, 6)">
          <v-subheader v-if="item.header" :key="item.header">{{ item.header }}</v-subheader>
          <v-divider v-else-if="item.divider" :key="index" :inset="item.inset"></v-divider>
          <v-list-item v-else :key="item.title" @click="">
            <v-list-item-avatar>
              <img :src="item.avatar">
            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-title v-html="item.title"></v-list-item-title>
              <v-list-item-subtitle v-html="item.subtitle"></v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </template>
      </v-list>
    </div>
    <div v-else class="text-center">  
      <v-progress-circular
        :size="100"
        color="primary"
        indeterminate
      ></v-progress-circular>
    </div>
  </div>
  <div v-else>
    <p>ログインしてください</p>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class UserNotification extends Vue {
  private isReady: boolean = false;
  private uid: string = '';
  private notifiactions: any[] = [];

  private items: any = [
  { header: 'Today' },
        { avatar: 'https://cdn.vuetifyjs.com/images/lists/1.jpg', title: 'Brunch this weekend?', subtitle: "<span class='font-weight-bold'>Ali Connors</span> &mdash; I'll be in your neighborhood doing errands this weekend. Do you want to hang out?" },
        { divider: true, inset: true },
        { avatar: 'https://cdn.vuetifyjs.com/images/lists/2.jpg', title: 'Summer BBQ <span class="grey--text text--lighten-1">4</span>', subtitle: "<span class='font-weight-bold'>to Alex, Scott, Jennifer</span> &mdash; Wish I could come, but I'm out of town this weekend." },
        { divider: true, inset: true },
        { avatar: 'https://cdn.vuetifyjs.com/images/lists/3.jpg', title: 'Oui oui', subtitle: "<span class='font-weight-bold'>Sandra Adams</span> &mdash; Do you have Paris recommendations? Have you ever been?" },

  ];

  private created(): void {
    const token = this.getToken();
    if (token !== null) {
      const claims = JSON.parse(atob(token.split('.')[1]));
      this.uid = claims.uid;
    }
    this.getNotifications();
  }

  private getToken(): any {
    return localStorage.getItem('token');
  }

  private getNotifications(): void {
    this.isReady = false;
    const url = 'api/notification/' + this.uid;
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      alert('server error');
      this.isReady = true;
      return [];
    }).then((json) => {
      this.notifiactions = [];
      this.notifiactions = json;
      this.isReady = true;
    });
  }

}
</script>

<style scoped>
</style>
