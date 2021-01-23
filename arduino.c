int sen0= A0;
int sen1= A1;
int sen2= A2;
int sen3= A3;
int sen4= A4;
int sen5= A5;
int sval= 0;

void setup() {
  pinMode(sen0, INPUT);
  pinMode(sen1, INPUT);
  pinMode(sen2, INPUT);
  pinMode(sen3, INPUT);
  pinMode(sen4, INPUT);
  pinMode(sen5, INPUT);
  
  Serial.begin(128000);
}
 
void loop() {
  sval=0;
  sval= analogRead(sen0);
  if(sval > 128){
Serial.print('a');
  Serial.println(sval);
  }
  sval=0;
  sval= analogRead(sen1);
  if(sval > 128){
Serial.print('b');
  Serial.println(sval);
  }
  sval=0;
  sval= analogRead(sen2);
  if(sval > 128){
Serial.print('c');
  Serial.println(sval);
  }
  sval=0;
  sval= analogRead(sen3);
  if(sval > 128){
Serial.print('d');
  Serial.println(sval);
  }
  sval=0;
  sval= analogRead(sen4);
  if(sval > 128){
Serial.print('e');
  Serial.println(sval);
  }
  sval=0;
  sval= analogRead(sen5);
  if(sval > 128){
Serial.print('f');
  Serial.println(sval);
  }
}