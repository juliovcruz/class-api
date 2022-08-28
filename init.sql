CREATE TABLE classes (
       id uuid primary key,
       name varchar not null,
       description varchar not null,
       total_hours integer not null,
       schedules varchar not null,
       teacher_id uuid not null,
       deleted boolean not null,
       created_at timestamp not null,
       updated_at timestamp not null
);

CREATE TABLE student_classes (
     id uuid primary key,
     class_id uuid not null,
     student_id uuid not null,
     created_at timestamp not null,
     updated_at timestamp not null,
     deleted_at timestamp null,
     CONSTRAINT fk_class FOREIGN KEY(class_id) REFERENCES classes(id)
);
