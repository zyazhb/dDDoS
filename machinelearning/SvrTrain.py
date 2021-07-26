from sklearn.svm import SVR  # SVM中的回归算法
import pandas as pd
from sklearn.model_selection import train_test_split
import matplotlib.pyplot as plt
import numpy as np
# 数据预处理，使得数据更加有效的被模型或者评估器识别
from sklearn import preprocessing
import joblib
 
# 获取数据
origin_data = pd.read_csv('test.csv')

# 处理时间数据
import time
timestamp=origin_data.iloc[:,2]
j=0
for i in timestamp:
    # 14/02/2018 08:31:01
    local=time.strptime(i,'%d/%m/%Y %H:%M:%S')
    origin_data.iloc[j,2]=time.mktime(local)
    j+=1
# print(origin_data.iloc[:,2])

#处理标签数据
# labelname=origin_data.iloc[:,79].unique() #['Benign','FTP-BruteForce']
labelname=origin_data.iloc[:,79]
for i in range(99):
    if((origin_data.iloc[i:,79]=='Benign').item):
        origin_data.iloc[i:,79]=0
    if((origin_data.iloc[i:,79]=='FTP-BruteForce').item):
        origin_data.iloc[i:,79]=1

X = origin_data.iloc[:,:].values
Y = origin_data.iloc[:,79].values
# print(type(Y))
# print(type(Y.values))

# 总特征  按照特征的重要性排序的所有特征
all_feature = [57, 55, 67, 75, 39, 56, 43, 29, 44, 19, 35, 20, 48, 16, 49, 34, 47, 63, 71, 66, 79, 68, 74, 72, 65, 77, 69, 31, 21, 25, 30, 27, 22, 1, 3, 4, 59, 60, 64, 62, 17, 45, 52, 41, 38, 9, 23, 10, 51, 24, 11, 12, 13, 18, 50, 40, 53, 14, 8, 6, 5, 7, 15, 32, 42, 58, 37, 54, 46, 36, 61, 33, 2, 28, 26, 73, 70, 78, 76, 0]
# 这里我们选取前三个特征
topN_feature = all_feature[:30]
print(topN_feature)
 
# 获取重要特征的数据
data_X = X[:,topN_feature]
 
# 将每个特征值归一化到一个固定范围
# 原始数据标准化，为了加速收敛
# 最小最大规范化对原始数据进行线性变换，变换到[0,1]区间
data_X = preprocessing.MinMaxScaler().fit_transform(data_X)
 
# 利用train_test_split 进行训练集和测试集进行分开
X_train,X_test,y_train,y_test  = train_test_split(data_X,Y,test_size=0.3)
 
# 通过多种模型预测
model_svr1 = SVR(kernel='rbf',C=50,max_iter=10000)
 
 
# 训练
# model_svr1.fit(data_X,Y)
model_svr1.fit(X_train,y_train)
 
# 得分
score = model_svr1.score(X_test,y_test)
print(score)