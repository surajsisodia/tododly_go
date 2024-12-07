select *
  from fusion.tasks;

select *
  from fusion.users;

select *
  from fusion.user_credentials;

DESCRIBE fusion.tasks;

drop table tasks;
drop table users;
drop table user_credentials;

commit;

ALTER TABLE tasks ADD deleted_at DATE;


alter table fusion.tasks drop column deleted  _at ;


create table tasks (
  task_id         NUMBER PRIMARY KEY,
  title           varchar2(100) not null,
  description     varchar2(255),
  created_at      date not null,
  created_by      varchar2(30) not null,
  last_updated_at date not null,
  last_updated_by varchar2(30) not null
);




-- Create the sequence
CREATE SEQUENCE task_id_s
START WITH 1
INCREMENT BY 1
NOCACHE;

-- Create the table
CREATE TABLE tasks (
    task_id NUMBER PRIMARY KEY,
    title VARCHAR2(100) NOT NULL,
    description VARCHAR2(255),
    created_at DATE NOT NULL,
    created_by VARCHAR2(30) NOT NULL,
    last_updated_at DATE NOT NULL,
    last_updated_by VARCHAR2(30) NOT NULL
);

-- Create the trigger
CREATE OR REPLACE TRIGGER trg_task_id
BEFORE INSERT ON tasks
FOR EACH ROW
WHEN (NEW.task_id IS NULL)
BEGIN
    SELECT task_id_s.NEXTVAL INTO :NEW.task_id FROM dual;
END;
/
