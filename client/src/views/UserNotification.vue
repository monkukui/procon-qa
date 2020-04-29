<template>
  <div v-if="uid">
    <h1>通知</h1>
    <div v-if="isReady">
      <v-list two-line>
        <div v-for="(item, index) in notifiactions">
          <UserNotificationPanel
            :ouid="item.ouid"
            :qid="item.qid"
            :type="item.type"
            :body="item.body"
            :watched="item.watched"
            :shouldDivide="index !== 0"
          />
        </div>
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
import UserNotificationPanel from '@/components/UserNotificationPanel.vue';

@Component({
  components: {
    UserNotificationPanel,
  },
})
export default class UserNotification extends Vue {
  private isReady: boolean = false;
  private uid: string = '';
  private notifiactions: any[] = [];

  private created(): void {
    const token = this.getToken();
    if (token !== null) {
      const claims = JSON.parse(atob(token.split('.')[1]));
      this.uid = claims.uid;
    }
    this.getNotifications();
    this.userUpdate();
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
      this.isReady = true;
      return [];
    }).then((json) => {
      this.notifiactions = [];
      this.notifiactions = json;
      this.isReady = true;
    });
  }

  // uid さんの NotficationFlag を false にする
  private userUpdate(): void {
    const url = '/api/notification/' + this.uid + '/0';
    const method = 'PUT';
    const headers = {
      'Authorization': `Bearer ${this.getToken()}`,
      'Content-Type': 'application/json; charset=UTF-8',
    };
    fetch(url, {method, headers});
  }


}
</script>

<style scoped>
</style>
