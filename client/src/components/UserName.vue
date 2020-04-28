<template>
  <span v-bind:class="color">
    <span v-if="isExist">
      <router-link 
        class="name"
        :to="{ name: 'userpage', query: { uid: this.uid }}"
      >
        {{ name }}
      </router-link>
    </span>
    <span v-else>
      {{ name }}
    </span>
  </span>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class UserName extends Vue {
  @Prop()
  private uid!: string;
  @Prop()
  private color!: string;

  private name: string = '';
  private isExist: boolean = false;


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
      if (this.name === '') {
        this.name = '退会済みのユーザ';
      } else {
        this.isExist = true;
      }
    });
  }
}
</script>

<style scoped>
.user-red {color:#FF0000;}
.user-orange {color:#FF8000;}
.user-yellow {color:#C0C000;}
.user-blue {color:#0000FF;}
.user-cyan {color:#00C0C0;}
.user-green {color:#008000;}
.user-brown {color:#804000;}
.user-gray {color:#808080;}
.user-unrated {color:#000000;}
.user-admin {color:#C000C0;}
.username>span {font-weight:bold;}
.name {
  font-weight: bold;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: rgb(66, 66, 66);
  text-decoration: none;
}
</style>
