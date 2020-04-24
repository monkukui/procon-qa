<template>
  <div> 
    <v-text-field>
    </v-text-field>
    <v-text-field>
      <v-icon slot="prepend" color="green">mdi-at</v-icon>
    </v-text-field>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';

@Component({
  components: {
  },
})

export default class UserUpdateForms extends Vue {

  private userName: string = '';
  private created(): void {
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.userName = claims.name;
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>
