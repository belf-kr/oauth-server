# 유틸
create database belf;
use belf;
show databases;
show tables;

# 데이터 조회
select *
from users;
select *
from kakao_talk_socials;

# 데이터 클리어
delete
from users;
delete
from kakao_talk_socials;

# 더미 데이터
insert into kakao_talk_socials (id, email, nick_name)
values (999, 'user01@test.com', '사용자1');
insert into users (email, password, name, connected_at, kakao_talk_socials_id)
values ('user01@test.com', sha2('user011!', 256), '사용자1', utc_timestamp(), 999);
insert into users (email, password, name, connected_at)
values ('user02@test.com', sha2('user021!', 256), '사용자2', utc_timestamp());

# 시드 데이터 삭제
delete
from users
where email = 'user01@test.com';

delete
from kakao_talk_socials
where id = 999;