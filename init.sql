CREATE TABLE classes (
       id uuid primary key,
       name varchar not null,
       description varchar not null,
       total_hours integer not null,
       schedules varchar not null,
       teacher_id uuid not null,
       created_at timestamp not null,
       updated_at timestamp not null,
       deleted_at timestamp null
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

create unique index student_class_student_class_unq on student_classes (class_id, student_id)
    where deleted_at is null;

create unique index classes_name_unq on classes(name)
    where deleted_at is null;