class Bob {

  constructor(str) {
    this._QUESTION = 'Sure.';
    this._YELL = 'Whoa, chill out!';
    this._SILENCE = 'Fine. Be that way!';
    this._OTHER = 'Whatever.';
  }

  hey(message) {
    this._str = message;
    return (this._isSilence() || this._isYell() || this._isQuestion() || this._OTHER);
  }

  _isQuestion() {
    if (this._str[this._str.length - 1] == '?') return this._QUESTION;
    return false;
  }

  _isYell() {
    let str = this._str;
    return (str === str.toUpperCase() && str.toLowerCase() != str ? this._YELL : false);
  }

  _isSilence() {
    return (!!(/^\s*$/.test(this._str)) ? this._SILENCE : false);
  }

}

module.exports = Bob;
/*
# Bob

Bob is a lackadaisical teenager. In conversation, his responses are
very limited.

Bob answers 'Sure.' if you ask him a question.

He answers 'Whoa, chill out!' if you yell at him.

He says 'Fine. Be that way!' if you address him without actually saying
anything.

He answers 'Whatever.' to anything else.

*/
