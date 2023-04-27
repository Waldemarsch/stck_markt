from airflow import DAG
from airflow.decorators import dag, task
from airflow.operators.python import PythonOperator
from airflow.providers.mysql.operators.mysql import MySqlOperator

from datetime import timedelta, datetime

default_args = {
    'owner': 'vladimir',
    'retries': 4,
    'retry_delay': timedelta(minutes=4)
}

def train():
    print(555555555555555555)

with DAG(dag_id='db_train',
     default_args=default_args,
     start_date=datetime(2023, 4, 23),
     schedule_interval='0 0 * * *',
    ) as dag:
    task1 = MySqlOperator(
        task_id='cr_table',
        mysql_conn_id="mysql_connection_train",
        sql="""
            create table if not exists popki (
                dt datetime,
                dag_id VARCHAR(40),
                primary key (dt, dag_id)
            )
        """
    )
    task2 = MySqlOperator(
        task_id='dupl_def',
        mysql_conn_id="mysql_connection_train",
        sql="""
            DELETE FROM popki
            WHERE dt = '{{ ts }}' AND dag_id = '{{ dag.dag_id }}'
        """
    )
    task3 = MySqlOperator(
        task_id='write_record',
        mysql_conn_id="mysql_connection_train",
        sql="""
            INSERT INTO popki
            VALUES('{{ ts }}', '{{ dag.dag_id }}')
        """
    )
    task1 >> task2 >> task3