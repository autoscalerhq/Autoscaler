
export class ApiClient {
  private sessionToken: string|undefined;
  public constructor(private baseUrl: string, sessionToken?: string) {
    this.sessionToken = sessionToken;

  }

  private async get(url: string) {
    return fetch(this.baseUrl + url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      ...(this.sessionToken ? {'Authorization': `Bearer ${this.sessionToken}`} : {})
      },
    });
  }

  private async post(url: string, body: BodyInit) {
    return fetch(this.baseUrl + url, {
      method: 'POST',
      body: body,
      headers: {
        'Content-Type': 'application/json',
      ...(this.sessionToken ? {'Authorization': `Bearer ${this.sessionToken}`} : {})
      }
    });
  }

  public async getComment() {
    return (await this.get('/comment')).text();
  }
}