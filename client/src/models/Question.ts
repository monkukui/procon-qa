export interface Question {

  id: number;
  uid?: number;
  tid?: number;

  title: string;
  body: string;
  date: string;
  state?: string;
  url?: string;
  completed: boolean;

  favoriteCount: number;
  answerCount: number;
  browseCount: number;
}
