PGDMP     9    3                x            daily_toolkit    12.3    12.3 "    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16910    daily_toolkit    DATABASE        CREATE DATABASE daily_toolkit WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';
    DROP DATABASE daily_toolkit;
                namhuynh    false                        3079    16911 	   uuid-ossp 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
    DROP EXTENSION "uuid-ossp";
                   false            �           0    0    EXTENSION "uuid-ossp"    COMMENT     W   COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';
                        false    2            �           1247    16996    meal_item_statuses    TYPE     P   CREATE TYPE public.meal_item_statuses AS ENUM (
    'empty',
    'completed'
);
 %   DROP TYPE public.meal_item_statuses;
       public          postgres    false            �            1259    16969    expense_aggregate    TABLE     �   CREATE TABLE public.expense_aggregate (
    _id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid,
    expense integer DEFAULT 0,
    created_at date DEFAULT CURRENT_DATE
);
 %   DROP TABLE public.expense_aggregate;
       public         heap    postgres    false    2            �            1259    16941    expense_categories    TABLE     �   CREATE TABLE public.expense_categories (
    _id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(30) NOT NULL,
    color character varying(7) NOT NULL,
    icon integer NOT NULL
);
 &   DROP TABLE public.expense_categories;
       public         heap    postgres    false    2            �            1259    16947    expenses    TABLE     �   CREATE TABLE public.expenses (
    _id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid,
    category_id uuid,
    amount integer DEFAULT 0,
    note text DEFAULT ''::text,
    created_at date DEFAULT CURRENT_DATE
);
    DROP TABLE public.expenses;
       public         heap    postgres    false    2            �            1259    17001 	   meal_item    TABLE     �   CREATE TABLE public.meal_item (
    _id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    meal_id uuid,
    kcal integer DEFAULT 0,
    status public.meal_item_statuses,
    created_at time without time zone DEFAULT CURRENT_TIME
);
    DROP TABLE public.meal_item;
       public         heap    postgres    false    2    657            �            1259    16982    meals    TABLE     �   CREATE TABLE public.meals (
    _id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid,
    kcal integer DEFAULT 0,
    created_at date DEFAULT CURRENT_DATE
);
    DROP TABLE public.meals;
       public         heap    postgres    false    2            �            1259    16929    user_settings    TABLE     �   CREATE TABLE public.user_settings (
    _id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid,
    max_kcal_per_day integer DEFAULT 10000
);
 !   DROP TABLE public.user_settings;
       public         heap    postgres    false    2            �            1259    16922    users    TABLE     �   CREATE TABLE public.users (
    _id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(30) NOT NULL,
    phone character varying(12) NOT NULL,
    created_at date DEFAULT CURRENT_DATE
);
    DROP TABLE public.users;
       public         heap    postgres    false    2            �          0    16969    expense_aggregate 
   TABLE DATA           N   COPY public.expense_aggregate (_id, user_id, expense, created_at) FROM stdin;
    public          postgres    false    207   V(       �          0    16941    expense_categories 
   TABLE DATA           D   COPY public.expense_categories (_id, name, color, icon) FROM stdin;
    public          postgres    false    205   s(       �          0    16947    expenses 
   TABLE DATA           W   COPY public.expenses (_id, user_id, category_id, amount, note, created_at) FROM stdin;
    public          postgres    false    206   �(       �          0    17001 	   meal_item 
   TABLE DATA           K   COPY public.meal_item (_id, meal_id, kcal, status, created_at) FROM stdin;
    public          postgres    false    209   _)       �          0    16982    meals 
   TABLE DATA           ?   COPY public.meals (_id, user_id, kcal, created_at) FROM stdin;
    public          postgres    false    208   |)       �          0    16929    user_settings 
   TABLE DATA           G   COPY public.user_settings (_id, user_id, max_kcal_per_day) FROM stdin;
    public          postgres    false    204   �)       �          0    16922    users 
   TABLE DATA           =   COPY public.users (_id, name, phone, created_at) FROM stdin;
    public          postgres    false    203   �)       �           2606    16976 (   expense_aggregate expense_aggregate_pkey 
   CONSTRAINT     g   ALTER TABLE ONLY public.expense_aggregate
    ADD CONSTRAINT expense_aggregate_pkey PRIMARY KEY (_id);
 R   ALTER TABLE ONLY public.expense_aggregate DROP CONSTRAINT expense_aggregate_pkey;
       public            postgres    false    207            �           2606    16946 *   expense_categories expense_categories_pkey 
   CONSTRAINT     i   ALTER TABLE ONLY public.expense_categories
    ADD CONSTRAINT expense_categories_pkey PRIMARY KEY (_id);
 T   ALTER TABLE ONLY public.expense_categories DROP CONSTRAINT expense_categories_pkey;
       public            postgres    false    205            �           2606    16958    expenses expenses_pkey 
   CONSTRAINT     U   ALTER TABLE ONLY public.expenses
    ADD CONSTRAINT expenses_pkey PRIMARY KEY (_id);
 @   ALTER TABLE ONLY public.expenses DROP CONSTRAINT expenses_pkey;
       public            postgres    false    206                       2606    17008    meal_item meal_item_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.meal_item
    ADD CONSTRAINT meal_item_pkey PRIMARY KEY (_id);
 B   ALTER TABLE ONLY public.meal_item DROP CONSTRAINT meal_item_pkey;
       public            postgres    false    209                        2606    16989    meals meals_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY public.meals
    ADD CONSTRAINT meals_pkey PRIMARY KEY (_id);
 :   ALTER TABLE ONLY public.meals DROP CONSTRAINT meals_pkey;
       public            postgres    false    208            �           2606    16935     user_settings user_settings_pkey 
   CONSTRAINT     _   ALTER TABLE ONLY public.user_settings
    ADD CONSTRAINT user_settings_pkey PRIMARY KEY (_id);
 J   ALTER TABLE ONLY public.user_settings DROP CONSTRAINT user_settings_pkey;
       public            postgres    false    204            �           2606    16928    users users_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (_id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    203                       2606    16977 0   expense_aggregate expense_aggregate_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.expense_aggregate
    ADD CONSTRAINT expense_aggregate_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(_id);
 Z   ALTER TABLE ONLY public.expense_aggregate DROP CONSTRAINT expense_aggregate_user_id_fkey;
       public          postgres    false    207    203    3062                       2606    16964 "   expenses expenses_category_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.expenses
    ADD CONSTRAINT expenses_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.expense_categories(_id);
 L   ALTER TABLE ONLY public.expenses DROP CONSTRAINT expenses_category_id_fkey;
       public          postgres    false    205    206    3066                       2606    16959    expenses expenses_user_id_fkey    FK CONSTRAINT     ~   ALTER TABLE ONLY public.expenses
    ADD CONSTRAINT expenses_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(_id);
 H   ALTER TABLE ONLY public.expenses DROP CONSTRAINT expenses_user_id_fkey;
       public          postgres    false    3062    206    203                       2606    17009     meal_item meal_item_meal_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.meal_item
    ADD CONSTRAINT meal_item_meal_id_fkey FOREIGN KEY (meal_id) REFERENCES public.meals(_id);
 J   ALTER TABLE ONLY public.meal_item DROP CONSTRAINT meal_item_meal_id_fkey;
       public          postgres    false    208    209    3072                       2606    16990    meals meals_user_id_fkey    FK CONSTRAINT     x   ALTER TABLE ONLY public.meals
    ADD CONSTRAINT meals_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(_id);
 B   ALTER TABLE ONLY public.meals DROP CONSTRAINT meals_user_id_fkey;
       public          postgres    false    208    203    3062                       2606    16936 (   user_settings user_settings_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.user_settings
    ADD CONSTRAINT user_settings_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(_id);
 R   ALTER TABLE ONLY public.user_settings DROP CONSTRAINT user_settings_user_id_fkey;
       public          postgres    false    203    204    3062            �      x������ � �      �   G   x�ȱ�0 �ְ�;��0��U�F��_�hV8(Ē4� �T�5+����;���0���Ƶ1n3"~Fr�      �   �   x�̱�0@�ڞ"�SdY��`K������ L@���/_�{�h�d3�ԃ�1�/[��'l"Z���p� k��d���$��R���uZ�@M�m��$Qūt?�f�����n��s$BB@��q�9�Z�)7      �      x������ � �      �      x������ � �      �      x������ � �      �   G   x����0�:^��y�K0MJ��-wduv-ꉥ3�:9�-w�Y�_�ҫ{�AD��Ԩ8��E�2��     